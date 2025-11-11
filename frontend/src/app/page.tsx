"use client"
import { useState } from "react"

export default function NovaAcademia() {
  interface AcademiaForm {
    nome: string
    endereco: string
    telefone: string
    preco: string
  }

  const [form, setForm] = useState<AcademiaForm>({
    nome: "",
    endereco: "",
    telefone: "",
    preco: ""
  })
  const [imagem, setImagem] = useState<File | null>(null)

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  const submit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    const data = new FormData()
    data.append("nome", form.nome)
    data.append("endereco", form.endereco)
    data.append("telefone", form.telefone)
    data.append("preco", form.preco)
    if (imagem) data.append("imagem", imagem)

    const res = await fetch("http://localhost:8081/academia", {
      method: "POST",
      body: data
    })

    const json = await res.json()
    console.log(json)
  }

  return (
    <form onSubmit={submit}>
      <input name="nome" onChange={handleChange} placeholder="Nome" />
      <input name="endereco" onChange={handleChange} placeholder="Endereço" />
      <input name="telefone" onChange={handleChange} placeholder="Telefone" />
      <input name="preco" onChange={handleChange} placeholder="Preço" />
      <input type="file" accept="image/*" onChange={e => setImagem(e.currentTarget.files?.[0] ?? null)} />
      <button type="submit">Enviar</button>
    </form>
  )
}
