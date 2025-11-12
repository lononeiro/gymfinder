"use client"

import { useState } from "react"
import { useRouter } from "next/navigation"

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8081"

export default function CadastroPage() {
  const [nome, setNome] = useState("")
  const [email, setEmail] = useState("")
  const [senha, setSenha] = useState("")
  const [admin, setAdmin] = useState(false)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState("")
  const [success, setSuccess] = useState("")
  const router = useRouter()

  async function handleCadastro(e: React.FormEvent) {
    e.preventDefault()
    setError("")
    setSuccess("")
    setLoading(true)

    try {
      const res = await fetch(`${API_URL}/usuario`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ nome, email, senha, admin }),
      })

      const data = await res.json()

      if (!res.ok) {
        setError(data.error || "Erro ao criar usuário.")
        setLoading(false)
        return
      }

      setSuccess("Usuário criado com sucesso!")
      setTimeout(() => router.push("/login"), 1500)
    } catch (err) {
      console.error(err)
      setError("Erro ao tentar se cadastrar.")
    }

    setLoading(false)
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-slate-50">
      <div className="bg-white p-8 rounded-2xl shadow-lg w-full max-w-md">
        <h1 className="text-3xl font-bold mb-6 text-center">Crie sua conta</h1>

        <form onSubmit={handleCadastro} className="flex flex-col gap-4">
          <div>
            <label className="block text-sm font-medium text-slate-600 mb-1">Nome</label>
            <input
              type="text"
              value={nome}
              onChange={(e) => setNome(e.target.value)}
              className="w-full border border-slate-300 rounded-lg p-2 focus:ring-2 focus:ring-blue-500 outline-none"
              placeholder="Digite seu nome completo"
              required
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-slate-600 mb-1">E-mail</label>
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="w-full border border-slate-300 rounded-lg p-2 focus:ring-2 focus:ring-blue-500 outline-none"
              placeholder="Digite seu e-mail"
              required
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-slate-600 mb-1">Senha</label>
            <input
              type="password"
              value={senha}
              onChange={(e) => setSenha(e.target.value)}
              className="w-full border border-slate-300 rounded-lg p-2 focus:ring-2 focus:ring-blue-500 outline-none"
              placeholder="Digite uma senha"
              required
            />
          </div>

          <label className="flex items-center gap-2 text-sm text-slate-700 mt-1">
            <input
              type="checkbox"
              checked={admin}
              onChange={(e) => setAdmin(e.target.checked)}
              className="w-4 h-4"
            />
            Usuário administrador
          </label>

          {error && <p className="text-red-500 text-sm mt-1">{error}</p>}
          {success && <p className="text-green-600 text-sm mt-1">{success}</p>}

          <button
            type="submit"
            disabled={loading}
            className={`mt-4 w-full rounded-lg p-3 font-semibold text-white transition ${
              loading ? "bg-slate-400" : "bg-blue-600 hover:bg-blue-700"
            }`}
          >
            {loading ? "Cadastrando..." : "Cadastrar"}
          </button>
        </form>

        <p className="text-sm text-center mt-4 text-slate-600">
          Já tem uma conta?{" "}
          <a href="/login" className="text-blue-600 hover:underline">
            Fazer login
          </a>
        </p>
      </div>
    </div>
  )
}
