const API_URL = process.env.NEXT_PUBLIC_API_URL || "https://gymfinder-1.onrender.com";

/**
 * Resolve URLs de imagens do backend para o formato correto
 * Lida com URLs do Filebase com ou sem protocolo, e arquivos locais
 */
export function resolveImageUrl(image?: string | null): string | null {
  if (!image) return null;

  const url = image.trim();

  // Se já tem protocolo, usa direto
  if (url.startsWith("http://") || url.startsWith("https://")) {
    return url;
  }

  // Se é URL do Filebase sem protocolo, adiciona https://
  if (url.includes("myfilebase.com") || url.includes("ipfs/")) {
    return `https://${url}`;
  }

  // Se é arquivo local no servidor
  return `${API_URL}/uploads/${url}`;
}

/**
 * Extrai URL de um objeto de imagem que pode vir em diferentes formatos
 */
export function extractImageUrl(item: any): string | null {
  if (!item) return null;

  // Se é string, retorna direto
  if (typeof item === "string") {
    return item;
  }

  // Se é objeto, procura por url ou nome_arquivo
  if (typeof item === "object") {
    return item.url || item.nome_arquivo || null;
  }

  return null;
}