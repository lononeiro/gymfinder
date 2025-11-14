"use client"

import React, { useState } from "react"

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8081"

export default function NovaAcademiaPage() {
  const [form, setForm] = useState({
    nome: "",
    endereco: "",
    telefone: "",
    preco: "",
    descricao: "",
  })
  const [imagens, setImagens] = useState<File[]>([])
  const [previewUrls, setPreviewUrls] = useState<string[]>([])
  const [loading, setLoading] = useState(false)
  const [mensagem, setMensagem] = useState("")

  function handleChange(e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  function handleFiles(e: React.ChangeEvent<HTMLInputElement>) {
    const files = e.target.files ? Array.from(e.target.files) : []
    setImagens(files)

    // gera previews
    const urls = files.map((f) => URL.createObjectURL(f))
    setPreviewUrls(urls)
  }

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    setMensagem("")
    setLoading(true)

    try {
      const token = localStorage.getItem("jwt") || ""
      const formData = new FormData()
      formData.append("nome", form.nome)
      formData.append("endereco", form.endereco)
      formData.append("telefone", form.telefone)
      formData.append("preco", form.preco)
      formData.append("descricao", form.descricao)

      imagens.forEach((img) => formData.append("imagens", img))

      const res =await fetch(`${API_URL}/academia`, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              Authorization: `Bearer ${token}`
            },
            body: JSON.stringify(form)
          })

      if (!res.ok) {
        const erroTexto = await res.text()
        console.log("ERRO DO BACKEND:", erroTexto)
        throw new Error("Erro ao criar academia: " + erroTexto)
        }

      setMensagem("✅ Academia criada com sucesso!")
      setForm({ nome: "", endereco: "", telefone: "", preco: "", descricao: "" })
      setImagens([])
      setPreviewUrls([])
    } catch (err) {
      console.error(err)
      setMensagem("❌ Erro ao criar academia.")
    }

    setLoading(false)
  }

  return (
    <div className="min-h-screen bg-slate-50 flex justify-center py-12 px-4">
      <div className="w-full max-w-3xl bg-white shadow-xl rounded-2xl p-8">
        <h1 className="text-3xl font-semibold text-center mb-8">Cadastrar Nova Academia</h1>

        <form onSubmit={handleSubmit} className="space-y-6">
          <div className="grid md:grid-cols-2 gap-4">
            <div>
              <label className="block font-medium text-slate-700">Nome</label>
              <input
                name="nome"
                value={form.nome}
                onChange={handleChange}
                required
                className="w-full border rounded-lg p-3 mt-1 focus:ring-2 focus:ring-black outline-none"
              />
            </div>

            <div>
              <label className="block font-medium text-slate-700">Telefone</label>
              <input
                name="telefone"
                value={form.telefone}
                onChange={handleChange}
                placeholder="(11) 99999-9999"
                className="w-full border rounded-lg p-3 mt-1 focus:ring-2 focus:ring-black outline-none"
              />
            </div>
          </div>

          <div>
            <label className="block font-medium text-slate-700">Endereço</label>
            <input
              name="endereco"
              value={form.endereco}
              onChange={handleChange}
              required
              className="w-full border rounded-lg p-3 mt-1 focus:ring-2 focus:ring-black outline-none"
            />
          </div>

          <div className="grid md:grid-cols-2 gap-4">
            <div>
              <label className="block font-medium text-slate-700">Preço</label>
              <input
                name="preco"
                value={form.preco}
                onChange={handleChange}
                placeholder="R$ 150,00"
                className="w-full border rounded-lg p-3 mt-1 focus:ring-2 focus:ring-black outline-none"
              />
            </div>
          </div>

          <div>
            <label className="block font-medium text-slate-700">Descrição</label>
            <textarea
              name="descricao"
              value={form.descricao}
              onChange={handleChange}
              rows={4}
              className="w-full border rounded-lg p-3 mt-1 focus:ring-2 focus:ring-black outline-none resize-none"
            />
          </div>

          <div>
            <label className="block font-medium text-slate-700">Imagens (você pode selecionar várias)</label>
            <input
              type="file"
              multiple
              accept="image/*"
              onChange={handleFiles}
              className="mt-2"
            />

            {previewUrls.length > 0 && (
              <div className="mt-4 flex gap-3 flex-wrap">
                {previewUrls.map((url, i) => (
                  <img
                    key={i}
                    src={url}
                    alt={`Preview ${i + 1}`}
                    className="w-28 h-28 object-cover rounded-lg shadow"
                  />
                ))}
              </div>
            )}
          </div>

          <button
            type="submit"
            disabled={loading}
            className="w-full bg-black text-white py-3 rounded-lg hover:bg-slate-800 transition disabled:opacity-50"
          >
            {loading ? "Enviando..." : "Criar Academia"}
          </button>

          {mensagem && (
            <p
              className={`text-center mt-4 font-medium ${
                mensagem.includes("✅") ? "text-green-600" : "text-red-600"
              }`}
            >
              {mensagem}
            </p>
          )}
        </form>
      </div>
    </div>
  )
}
