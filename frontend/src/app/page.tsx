"use client";
import { useEffect, useState } from "react";
import { api } from "@/services/api";

interface Academia {
  id: number;
  nome: string;
  endereco: string;
  preco: string;
  descricao: string;
}

export default function Home() {
  const [academias, setAcademias] = useState<Academia[]>([]);

  useEffect(() => {
    api.get("/academias")
      .then((res) => setAcademias(res.data.academias))
      .catch((err) => console.error(err));
  }, []);

  return (
    <main className="p-8">
      <h1 className="text-3xl font-bold mb-6">Academias</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {academias.map((a) => (
          <div key={a.id} className="bg-gray-800 p-4 rounded-xl shadow">
            <h2 className="text-xl font-semibold">{a.nome}</h2>
            <p>{a.endereco}</p>
            <p>{a.preco}</p>
            <p className="text-sm text-gray-400 mt-2">{a.descricao}</p>
          </div>
        ))}
      </div>
    </main>
  );
}