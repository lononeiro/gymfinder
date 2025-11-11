// ...existing code...
"use client"

import { useEffect, useState } from "react"

export default function AcademiasPage() {
  const [academias, setAcademias] = useState([])

  async function fetchAcademias() {
    const res = await fetch("http://localhost:8081/academias", { cache: "no-store" })
    const data = await res.json()
    setAcademias(data)
  }

  useEffect(() => {
    fetchAcademias()
  }, [])

  return (
    <div style={{ padding: "20px", display: "flex", flexWrap: "wrap", gap: "20px" }}>
      {academias.map((item: any) => {
        // Suporta os formatos retornados pelo backend:
        // - item.imagem como string (nome do arquivo)
        // - item.imagem como objeto com nome_arquivo
        // - fallback para item.imagens[0]?.url
        const imageName =
          (typeof item.imagem === "string" ? item.imagem : item.imagem?.nome_arquivo) ??
          item.imagens?.[0]?.url

        return (
          <div key={item.id} style={{ width: "300px", border: "1px solid #ddd", borderRadius: "10px", overflow: "hidden" }}>
            {imageName ? (
              <img
                src={`http://localhost:8081/uploads/${imageName}`}
                style={{ width: "100%", height: "200px", objectFit: "cover" }}
                alt={item.nome ?? "Academia"}
              />
            ) : (
              <div style={{ width: "100%", height: "200px", background: "#f0f0f0" }} />
            )}
            <div style={{ padding: "10px" }}>
              <h2 style={{ margin: 0 }}>{item.nome}</h2>
              <p style={{ marginTop: "5px" }}>{item.endereco}</p>
              <p style={{ marginTop: "5px" }}>Preço: {item.preco}</p>
            </div>
          </div>
        )
      })}
    </div>
  )
}
// ...existing code...