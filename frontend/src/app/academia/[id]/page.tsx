"use client";

import React, { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import CommentForm from "@/components/CommentForm";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "https://gymfinder-1.onrender.com";

// Função para normalizar URLs de imagem
function normalizeImageUrl(url?: string | null): string | null {
  if (!url) return null;
  
  // Remove qualquer prefixo incorreto que possa estar vindo
  if (url.includes("gymfinder-1.onrender.com/uploads/https://")) {
    return url.replace("gymfinder-1.onrender.com/uploads/https://", "");
  }
  
  // Se já começa com http ou https, retorna como está
  if (url.startsWith("http://") || url.startsWith("https://")) {
    return url;
  }
  
  // Se parece ser uma URL do Filebase mas sem protocolo, adiciona https://
  if (url.includes("future-coffee-galliform.myfilebase.com")) {
    return `https://${url}`;
  }
  
  // Para imagens locais, usa o endpoint correto
  return url;
}
''
export default function AcademiaDetalhePage() {
  const params = useParams();
  const id = params?.id;

  const [academia, setAcademia] = useState<any>(null);
  const [comentarios, setComentarios] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);

  // Normaliza o valor que pode ser: string (nome do arquivo ou URL) ou objeto { url, nome_arquivo }
  function normalizeImage(item: any): string | null {
    if (!item) return null;

    // extrai o campo que pode ser string ou objeto
    let value: string | undefined | null;
    if (typeof item === "string") {
      value = item;
    } else if (typeof item === "object") {
      value = item.url || item.nome_arquivo || null;
    }

    if (!value) return null;

    return normalizeImageUrl(value);
  }

  useEffect(() => {
    if (!id) return;

    setLoading(true);
    Promise.all([
      fetch(`${API_URL}/academia/${id}`).then((r) => r.json()),
      fetch(`${API_URL}/academia/${id}/comentarios`).then((r) => r.json()).catch(() => []),
    ])
      .then(([acadData, comms]) => {
        setAcademia(acadData);
        setComentarios(Array.isArray(comms) ? comms : []);
      })
      .catch((err) => {
        console.error("Erro ao carregar academia:", err);
        setAcademia(null);
        setComentarios([]);
      })
      .finally(() => setLoading(false));
  }, [id]);

  if (loading) return <div className="p-6">Carregando...</div>;
  if (!academia) return <div className="p-6">Academia não encontrada.</div>;

  // monta array de imagens normalizadas
  const imagens =
    academia.imagens && academia.imagens.length > 0
      ? academia.imagens
          .map((img: any) => normalizeImage(img))
          .filter((u: any) => !!u)
      : [];

  console.log("Imagens normalizadas:", imagens);

  return (
    <div className="max-w-5xl mx-auto px-4 py-8">
      <div className="rounded-xl shadow bg-white overflow-hidden">
        {imagens.length > 0 ? (
          <div className="flex overflow-x-auto space-x-3 p-2 scrollbar-thin">
            {imagens.map((src: string, i: number) => (
              <img
                key={i}
                src={src}
                alt={`${academia.nome} - ${i + 1}`}
                className="w-full max-w-[600px] h-auto object-cover rounded"
                loading="lazy"
                onError={(e) => {
                  console.error(`Erro ao carregar imagem: ${src}`);
                  // Esconde a imagem com erro
                  (e.target as HTMLImageElement).style.display = 'none';
                }}
              />
            ))}
          </div>
        ) : (
          <div className="w-full h-48 flex items-center justify-center bg-slate-100">
            Sem imagens
          </div>
        )}

        <div className="p-6">
          <h1 className="text-2xl font-bold">{academia.nome}</h1>
          <p className="mt-1 text-sm text-slate-600">{academia.endereco}</p>
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
          <CommentForm
            academiaId={academia.id}
            onSuccess={() => {
              fetch(`${API_URL}/academias/${id}/comentarios`).then((r) => r.json()).then((comms) => {
                setComentarios(Array.isArray(comms) ? comms : []);
              });
            }}
          />

          <div className="mt-6 space-y-3">
            {comentarios.map((c: any, idx: number) => (
              <div key={c.id ?? idx} className="bg-white rounded-xl shadow p-4">
                <p className="text-slate-800">{c.texto}</p>
                <p className="text-sm text-slate-500 mt-1">Por {c.usuario_nome}</p>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}