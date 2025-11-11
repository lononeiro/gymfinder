"use client"

import React, { useEffect, useState } from "react"
import { useParams } from "next/navigation"

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8081"

export default function AcademiaDetalhePage() {
  const params = useParams()
  const id = params?.id

  const [academia, setAcademia] = useState<any>(null)
  const [comentarios, setComentarios] = useState<any[]>([])
  const [loading, setLoading] = useState(true)
  const [novoComentario, setNovoComentario] = useState("")
  const [sending, setSending] = useState(false)

  async function fetchData() {
    try {
      const aca = await fetch(`${API_URL}/academias`) // backend não tem GET /academia/{id}, então buscamos todos
      const lista = await aca.json()
      const selecionada = lista.find((x: any) => String(x.id) === String(id))
      setAcademia(selecionada || null)

      const com = await fetch(`${API_URL}/academia/${id}/comentario`)
      const comData = await com.json()
      setComentarios(comData.comentarios || [])
    } catch (e) {
      console.error(e)
    }
    setLoading(false)
  }

  useEffect(() => {
    if (id) fetchData()
  }, [id])

  async function enviarComentario() {
    if (!novoComentario.trim()) return
    setSending(true)

    try {
      const token = localStorage.getItem("jwt") || ""
      await fetch(`${API_URL}/academia/${id}/comentario`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ texto: novoComentario }),
      })

      setNovoComentario("")
      fetchData()
    } catch (e) {
      console.error(e)
    }

    setSending(false)
  }

  if (loading) return <div className="p-6">Carregando...</div>
  if (!academia) return <div className="p-6">Academia não encontrada.</div>

  const imgName =
    (typeof academia.imagem === "string" ? academia.imagem : academia.imagem?.nome_arquivo) ??
    academia.imagens?.[0]?.url

  const imgUrl = imgName ? `${API_URL}/uploads/${imgName}` : null

  return (
    <div className="max-w-5xl mx-auto px-4 py-8">
      <div className="rounded-xl shadow bg-white overflow-hidden">
        {imgUrl ? (
          <img src={imgUrl} className="w-full h-72 object-cover" />
        ) : (
          <div className="w-full h-72 bg-slate-200 flex items-center justify-center">Sem imagem</div>
        )}

        <div className="p-6">
          <h1 className="text-3xl font-semibold">{academia.nome}</h1>
          <p className="mt-2 text-slate-700">{academia.endereco}</p>
          <p className="mt-3 font-medium">Preço: {academia.preco ?? "—"}</p>

          {academia.descricao && (
            <p className="mt-4 text-slate-600 whitespace-pre-line">{academia.descricao}</p>
          )}
        </div>
      </div>

      {/* Comentários */}
      <div className="mt-10">
        <h2 className="text-2xl font-semibold">Comentários</h2>

        <div className="mt-4 bg-white rounded-xl shadow p-4">
          <textarea
            className="w-full border rounded p-3 h-28 outline-none"
            placeholder="Escreva um comentário..."
            value={novoComentario}
            onChange={(e) => setNovoComentario(e.target.value)}
          />

          <button
            onClick={enviarComentario}
            disabled={sending}
            className="mt-3 bg-black text-white px-4 py-2 rounded hover:bg-slate-800 transition disabled:opacity-50"
          >
            {sending ? "Enviando..." : "Enviar"}
          </button>
        </div>

        <div className="mt-6 space-y-4">
          {comentarios.length === 0 && <p className="text-slate-500">Nenhum comentário ainda.</p>}

          {comentarios.map((c) => (
            <div key={c.id} className="bg-white rounded-xl shadow p-4">
              <p className="text-slate-800">{c.texto}</p>
              <p className="text-sm text-slate-500 mt-1">Por {c.usuario_nome || c.usuario?.nome || "Usuário"}</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}
