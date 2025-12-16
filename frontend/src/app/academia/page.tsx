"use client"

import React, { useEffect, useRef, useState } from "react"
import Link from "next/link"

const API_URL = process.env.NEXT_PUBLIC_API_URL || "https://gymfinder-1.onrender.com"

type Academia = {
  id: number
  nome: string
  endereco: string
  telefone?: string
  preco?: string
  descricao?: string
  imagem_principal?: string | null
  imagens?: { id: number; url: string }[]
}

// Função para normalizar URLs de imagem
function normalizeImageUrl(url?: string | null): string | null {
  if (!url) return null
  
  // Remove qualquer prefixo incorreto que possa estar vindo
  if (url.includes("gymfinder-1.onrender.com/uploads/https://")) {
    return url.replace("gymfinder-1.onrender.com/uploads/https://", "https://")
  }
  
  // Se já começa com http ou https, retorna como está
  if (url.startsWith("http://") || url.startsWith("https://")) {
    return url
  }
  
  // Se parece ser uma URL do Filebase mas sem protocolo, adiciona https://
  if (url.includes("future-coffee-galliform.myfilebase.com")) {
    return `https://${url}`
  }
  
  // Para imagens locais, usa o endpoint correto
  return `${API_URL}/uploads/${url}`
}

export default function AcademiasPage() {
  const [academias, setAcademias] = useState<Academia[]>([])
  const [currentSlide, setCurrentSlide] = useState(0)
  const autoPlayRef = useRef<number | null>(null)
  const cardRefs = useRef<Record<number, HTMLDivElement | null>>({})

  async function fetchAcademias() {
    try {
      const res = await fetch(`${API_URL}/academias`, { cache: "no-store" })
      const data = await res.json()
      const parsed = Array.isArray(data) ? data : []
      setAcademias(parsed)
      console.log("Academias fetched:", parsed)
    } catch (err) {
      console.error("Erro ao buscar academias", err)
      setAcademias([])
    }
  }

  useEffect(() => {
    fetchAcademias()
  }, [])

  // ---------------- CAROUSEL ----------------
  const slides = academias.length
    ? academias.map((item) => ({
        id: item.id,
        title: item.nome,
        subtitle: item.endereco,
        image: normalizeImageUrl(item.imagem_principal),
      }))
    : [
        {
          id: "placeholder-1",
          title: "Bem-vindo",
          subtitle: "Encontre sua academia",
          image: null,
        },
        {
          id: "placeholder-2",
          title: "Treine hoje",
          subtitle: "Procure perto de você",
          image: null,
        },
      ]

  useEffect(() => {
    stopAutoPlay()
    autoPlayRef.current = window.setInterval(() => {
      setCurrentSlide((s) => (s + 1) % slides.length)
    }, 4000)
    return () => stopAutoPlay()
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

  // ------------- INTERSECTION OBSERVER -------------
  useEffect(() => {
    const obs = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          const el = entry.target as HTMLDivElement
          if (entry.isIntersecting) el.classList.add("is-visible")
        })
      },
      { threshold: 0.12 }
    )

    Object.values(cardRefs.current).forEach((el) => {
      if (el) obs.observe(el)
    })

    return () => obs.disconnect()
  }, [academias])

  return (
    <div className="w-full min-h-screen bg-slate-50 text-slate-900">
      {/* ---------------- BIG CAROUSEL ---------------- */}
      <section className="w-full relative">
        <div className="w-full h-[56vh] md:h-[60vh] lg:h-[68vh] overflow-hidden relative">
          <div
            className="absolute inset-0 flex transition-transform duration-700 ease-out"
            style={{ transform: `translateX(-${currentSlide * 100}%)` }}
          >
            {slides.map((s, i) => (
              <div key={s.id} className="w-full flex-shrink-0 relative">
                {s.image ? (
                  <img
                    src={s.image}
                    alt={s.title}
                    className="w-full h-[56vh] md:h-[60vh] lg:h-[68vh] object-cover"
                    onError={(e) => {
                      const target = e.target as HTMLImageElement
                      target.src = "/placeholder.png"
                    }}
                  />
                ) : (
                  <div className="w-full h-[56vh] md:h-[60vh] lg:h-[68vh] flex items-center justify-center bg-gradient-to-r from-slate-300 via-slate-200 to-slate-300">
                    <div className="text-center">
                      <h1 className="text-3xl md:text-5xl font-semibold">
                        {s.title}
                      </h1>
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

          <button
            onClick={prev}
            className="absolute left-4 top-1/2 -translate-y-1/2 bg-white/80 shadow rounded-full p-3"
          >
            ‹
          </button>
          <button
            onClick={next}
            className="absolute right-4 top-1/2 -translate-y-1/2 bg-white/80 shadow rounded-full p-3"
          >
            ›
          </button>
        </div>
      </section>

      {/* ---------------- GRID ---------------- */}
      <section className="max-w-6xl mx-auto px-4 pb-16">
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          {academias.map((item) => {
            const imageUrl = normalizeImageUrl(item.imagem_principal)

            return (
              <Link href={`/academia/${item.id}`} key={item.id}>
                <div
                  ref={(el) => {
                    cardRefs.current[item.id] = el
                  }}
                  className="rounded-xl overflow-hidden bg-white shadow-sm border border-slate-100 transform opacity-0 translate-y-6 transition-all duration-600 ease-out hover:shadow-md cursor-pointer"
                >
                  {imageUrl ? (
                    <img
                      src={imageUrl}
                      alt={item.nome}
                      className="w-full h-48 object-cover"
                      onError={(e) => {
                        const target = e.target as HTMLImageElement
                        target.src = "/placeholder.png"
                      }}
                    />
                  ) : (
                    <div className="w-full h-48 flex items-center justify-center bg-slate-100">
                      Sem imagem
                    </div>
                  )}

                  <div className="p-4">
                    <h3 className="text-lg font-medium">{item.nome}</h3>
                    <p className="mt-1 text-sm text-slate-600">
                      {item.endereco}
                    </p>
                    <p className="mt-2 font-semibold">
                      Preço: {item.preco ?? "—"}
                    </p>
                  </div>
                </div>
              </Link>
            )
          })}
        </div>
      </section>

      <style jsx>{`
        .is-visible {
          opacity: 1 !important;
          transform: translateY(0) !important;
        }
      `}</style>
    </div>
  )
}