"use client"

import React, { useState } from "react"
import { getToken, isAuthenticated } from "@/lib/auth"

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8081"

export default function CommentForm({
  academiaId,
  onSuccess,
}: {
  academiaId: number | string
  onSuccess?: () => void
}) {
  const [texto, setTexto] = useState("")
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [success, setSuccess] = useState<string | null>(null)

  if (!isAuthenticated()) {
    return (
      <div className="p-4 bg-yellow-50 rounded">
        <p className="text-sm">Você precisa estar <a href="/login" className="font-semibold underline">logado</a> para comentar.</p>
      </div>
    )
  }

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    setError(null)
    setSuccess(null)
    setLoading(true)

    try {
      const token = getToken() || ""
      const res = await fetch(`${API_URL}/academia/${academiaId}/comentario`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ texto }),
      })

      if (!res.ok) {
        const txt = await res.text()
        throw new Error(txt || `Erro ${res.status}`)
      }

      setSuccess("Comentário enviado")
      setTexto("")
      if (onSuccess) onSuccess()
    } catch (err) {
      console.error(err)
      setError("Erro ao enviar comentário")
    } finally {
      setLoading(false)
    }
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-2">
      <textarea
        value={texto}
        onChange={(e) => setTexto(e.target.value)}
        placeholder="Escreva seu comentário..."
        className="w-full border rounded p-2"
        rows={3}
        required
      />
      <div>
        <button disabled={loading} className="px-4 py-2 bg-black text-white rounded">
          {loading ? "Enviando..." : "Comentar"}
        </button>
      </div>
      {error && <p className="text-red-600">{error}</p>}
      {success && <p className="text-green-600">{success}</p>}
    </form>
  )
}
