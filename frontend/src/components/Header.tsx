"use client";

import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export default function Header() {
  const pathname = usePathname();
  const router = useRouter();
  const [usuarioNome, setUsuarioNome] = useState<string | null>(null);

  useEffect(() => {
    if (typeof window !== "undefined") {
      const nome = localStorage.getItem("usuario_nome");
      setUsuarioNome(nome);
    }
  }, []);

  function handleLogout() {
    localStorage.clear();
    router.push("/login");
  }

  return  (
    <header className="bg-white/80 backdrop-blur-md border-b border-slate-200 shadow-sm sticky top-0 z-50">
      <div className="max-w-6xl mx-auto flex justify-between items-center px-4 py-3">
        {/* LOGO / HOME */}
        <Link href="/" className="text-2xl font-bold text-yellow-600">
          GymFinder
        </Link>

        {/* LINKS */}
        <nav className="flex gap-4 items-center">
          <Link
            href="/academia/nova-academia"
            className={`hover:text-yellow-600 transition ${
              pathname === "/nova-academia" ? "text-yellow-600 font-semibold" : ""
            }`}
          >
            Criar Academia
          </Link>

          {!usuarioNome ? (
            <Link
              href="/login"
              className={`bg-yellow-500 hover:bg-yellow-600 text-white px-4 py-2 rounded-lg transition`}
            >
              Login
            </Link>
          ) : (
            <div className="flex items-center gap-3">
              <span className="text-slate-700 font-medium">
              </span>
              <button
                onClick={handleLogout}
                className="text-sm text-red-500 hover:text-red-600 transition, mouseover:underline"
              >
                Sair
              </button>
            </div>
          )}
        </nav>
      </div>
    </header>
  );
}

