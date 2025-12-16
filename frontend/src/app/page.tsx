"use client"

import React, { useEffect, useRef, useState } from "react"
import Link from "next/link"

const API_URL =
  process.env.NEXT_PUBLIC_API_URL || "https://gymfinder-1.onrender.com"

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

/**
 * Normaliza qualquer formato de URL de imagem vindo do backend
 */
function normalizeImageUrl(url?: string | null): string | null {
  if (!url) return null

  let clean = url.trim()

  // Remove prefixo errado: /uploads/https://...
  if (clean.includes("/uploads/https://")) {
    clean = clean.split("/uploads/")[1]
  }

  // Se já for URL absoluta válida, retorna direto
  if (clean.startsWith("http://") || clean.startsWith("https://")) {
    return clean
  }

  // Caso seja apenas nome de arquivo local
  return `${API_URL.replace(/\/$/, "")}/uploads/${clean.replace(/^\/+/, "")}`
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
      setAcademias(Array.isArray(data) ? data : [])
    } catch (err) {
      console.error("Erro ao buscar academias:", err)
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

  // ------------- INTERSECTION OBSERVER -------------
  useEffect(() => {
    const obs = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            ;(entry.target as HTMLDivElement).classList.add("is-visible")
          }
        })
      },
      { threshold: 0.12 }
    )

    Object.values(cardRefs.current).forEach((el) => el && obs.observe(el))
    return () => obs.disconnect()
  }, [academias])

  return (
    <div className="w-full min-h-screen bg-slate-50 text-slate-900">
      {/* ---------------- CAROUSEL ---------------- */}
      <section className="w-full relative">
        <div className="w-full h-[56vh] md:h-[60vh] lg:h-[68vh] overflow-hidden relative">
          <div
            className="absolute inset-0 flex transition-transform duration-700 ease-out"
            style={{ transform: `translateX(-${currentSlide * 100}%)` }}
          >
            {slides.map((s) => (
              <div key={s.id} className="w-full flex-shrink-0 relative">
                {s.image ? (
                  <img
                    src={s.image}
                    alt={s.title}
                    className="w-full h-full object-cover"
                    loading="lazy"
                    referrerPolicy="no-referrer"
                  />
                ) : (
                  <div className="w-full h-full flex items-center justify-center bg-slate-200">
                    <h1 className="text-3xl font-semibold">{s.title}</h1>
                  </div>
                )}
              </div>
            ))}
          </div>

          <button
            onClick={prev}
            className="absolute left-4 top-1/2 -translate-y-1/2 bg-white/80 rounded-full p-3"
          >
            ‹
          </button>
          <button
            onClick={next}
            className="absolute right-4 top-1/2 -translate-y-1/2 bg-white/80 rounded-full p-3"
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
                    if (el) cardRefs.current[item.id] = el
                  }}
                  className="rounded-xl overflow-hidden bg-white shadow-sm border opacity-0 translate-y-6 transition-all duration-500 hover:shadow-md cursor-pointer"
                >
                  {imageUrl ? (
                    <img
                      src={imageUrl.replace("https://gymfinder-1.onrender.com/uploads/", "")}
                      alt={item.nome}
                      className="w-full h-48 object-cover"
                      loading="lazy"
                      referrerPolicy="no-referrer"
                    />
                  ) : (
                    <div className="w-full h-48 flex items-center justify-center bg-slate-100">
                      Sem imagem
                    </div>
                  )}

                  <div className="p-4">
                    <h3 className="text-lg font-medium">{item.nome}</h3>
                    <p className="text-sm text-slate-600">{item.endereco}</p>
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
