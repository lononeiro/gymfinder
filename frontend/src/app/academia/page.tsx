"use client"

import React, { useEffect, useRef, useState } from "react"
import Link from "next/link"

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8081"

export default function AcademiasPage() {
  const [academias, setAcademias] = useState<any[]>([])
  const [currentSlide, setCurrentSlide] = useState(0)
  const autoPlayRef = useRef<number | null>(null)
  const cardRefs = useRef<Array<HTMLDivElement | null>>([])

  async function fetchAcademias() {
    try {
      const res = await fetch(`${API_URL}/academias`, { cache: "no-store" })
      const data = await res.json()
      setAcademias(Array.isArray(data) ? data : [])
    } catch (err) {
      console.error("Erro ao buscar academias", err)
      setAcademias([])
    }
  }

  useEffect(() => {
    fetchAcademias()
  }, [])

  // --- Carousel logic ---
  const slides = academias.length
    ? academias.map((item) => {
        const imageName =
          (typeof item.imagem === "string" ? item.imagem : item.imagem?.nome_arquivo) ??
          item.imagens?.[0]?.url
        return {
          id: item.id,
          title: item.nome,
          subtitle: item.endereco,
          image: imageName ? `${API_URL}/uploads/${imageName}` : null,
        }
      })
    : [
        { id: "placeholder-1", title: "Bem-vindo", subtitle: "Encontre sua academia", image: null },
        { id: "placeholder-2", title: "Treine hoje", subtitle: "Procure perto de você", image: null },
      ]

  useEffect(() => {
    // autoplay
    stopAutoPlay()
    autoPlayRef.current = window.setInterval(() => {
      setCurrentSlide((s) => (s + 1) % slides.length)
    }, 4000)
    return () => stopAutoPlay()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [slides.length])

  function stopAutoPlay() {
    if (autoPlayRef.current) {
      clearInterval(autoPlayRef.current)
      autoPlayRef.current = null
    }
  }

  function prev() {
    stopAutoPlay()
    setCurrentSlide((s) => (s - 1 + slides.length) % slides.length)
  }

  function next() {
    stopAutoPlay()
    setCurrentSlide((s) => (s + 1) % slides.length)
  }

  // --- Scroll animation for cards ---
  useEffect(() => {
    const obs = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          const el = entry.target as HTMLDivElement
          if (entry.isIntersecting) {
            el.classList.add("is-visible")
          }
        })
      },
      { threshold: 0.12 }
    )

    cardRefs.current.forEach((el) => {
      if (el) obs.observe(el)
    })

    return () => obs.disconnect()
  }, [academias])

  return (
    <div className="w-full min-h-screen bg-slate-50 text-slate-900">
      {/* Big carousel */}
      <section className="w-full relative">
        <div className="w-full h-[56vh] md:h-[60vh] lg:h-[68vh] overflow-hidden relative">
          <div className="absolute inset-0 flex transition-transform duration-700 ease-out" style={{ transform: `translateX(-${currentSlide * 100}%)` }}>
            {slides.map((s, i) => (
              <div key={s.id} className="w-full flex-shrink-0 relative">
                {s.image ? (
                  <img src={s.image} alt={s.title} className="w-full h-[56vh] md:h-[60vh] lg:h-[68vh] object-cover" />
                ) : (
                  <div className="w-full h-[56vh] md:h-[60vh] lg:h-[68vh] flex items-center justify-center bg-gradient-to-r from-slate-300 via-slate-200 to-slate-300">
                    <div className="text-center">
                      <h1 className="text-3xl md:text-5xl font-semibold">{s.title}</h1>
                      <p className="mt-2 text-lg md:text-xl">{s.subtitle}</p>
                    </div>
                  </div>
                )}
                <div className="absolute left-6 bottom-8 bg-white/60 backdrop-blur-md rounded-md px-4 py-2">
                  <h3 className="text-lg font-medium">{s.title}</h3>
                  <p className="text-sm">{s.subtitle}</p>
                </div>
              </div>
            ))}
          </div>

          {/* Prev / Next */}
          <button onClick={prev} aria-label="Anterior" className="absolute left-4 top-1/2 -translate-y-1/2 bg-white/80 shadow rounded-full p-3 hover:scale-105 transition">
            ‹
          </button>
          <button onClick={next} aria-label="Próximo" className="absolute right-4 top-1/2 -translate-y-1/2 bg-white/80 shadow rounded-full p-3 hover:scale-105 transition">
            ›
          </button>

          {/* Dots */}
          <div className="absolute left-1/2 -translate-x-1/2 bottom-6 flex gap-2">
            {slides.map((_, i) => (
              <button
                key={i}
                onClick={() => { stopAutoPlay(); setCurrentSlide(i) }}
                className={`w-3 h-3 rounded-full transition-all ${i === currentSlide ? 'scale-110 bg-white' : 'bg-white/60'}`}
                aria-label={`Slide ${i + 1}`}
              />
            ))}
          </div>
        </div>
      </section>

      {/* Small animated intro when user scrolls down */}
      <section className="max-w-6xl mx-auto px-4 py-8 md:py-12">
        <div className="bg-white rounded-2xl shadow p-6 md:p-10 transform -translate-y-10 md:-translate-y-12 opacity-0 animate-slide-up-once">
          <h2 className="text-2xl md:text-3xl font-semibold">Academias perto de você</h2>
          <p className="mt-2 text-sm text-slate-600">Role um pouco para ver as academias — os cards vão aparecer com uma pequena animação.</p>
        </div>
      </section>

      {/* Cards grid */}
      <section className="max-w-6xl mx-auto px-4 pb-16">
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          {academias.map((item, idx) => {
            const imageName =
              (typeof item.imagem === "string" ? item.imagem : item.imagem?.nome_arquivo) ?? item.imagens?.[0]?.url
            const imageUrl = imageName ? `${API_URL}/uploads/${imageName}` : null

            return (
              <Link href={`/academia/${item.id ?? idx}`} key={item.id ?? idx} className="block">
                {/* o ref precisa estar na div interna, não no Link */}
                <div
                  ref={(el) => { cardRefs.current[idx] = el }}
                  className="rounded-xl overflow-hidden bg-white shadow-sm border border-slate-100 transform opacity-0 translate-y-6 transition-all duration-600 ease-out hover:shadow-md focus:shadow-md focus:outline-none cursor-pointer"
                  style={{ transitionTimingFunction: 'cubic-bezier(.2,.9,.2,1)' }}
                  tabIndex={0}
                >
                  {imageUrl ? (
                    <img src={imageUrl} alt={item.nome} className="w-full h-48 object-cover" />
                  ) : (
                    <div className="w-full h-48 flex items-center justify-center bg-slate-100">Sem imagem</div>
                  )}

                  <div className="p-4">
                    <h3 className="text-lg font-medium">{item.nome}</h3>
                    <p className="mt-1 text-sm text-slate-600">{item.endereco}</p>
                    <p className="mt-2 font-semibold">Preço: {item.preco ?? '—'}</p>
                  </div>
                </div>
              </Link>
            )
          })}
        </div>
      </section>

      <style jsx>{`
        /* small helper animation (one-time) for the intro card */
        @keyframes slideUpOnce {
          from { transform: translateY(8px); opacity: 0 }
          to { transform: translateY(0); opacity: 1 }
        }
        .animate-slide-up-once {
          animation: slideUpOnce 700ms ease-out forwards;
        }

        /* when card gets into view we add .is-visible (IntersectionObserver) */
        .is-visible {
          opacity: 1 !important;
          transform: translateY(0) !important;
        }
      `}</style>
    </div>
  )
}
