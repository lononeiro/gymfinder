export function getToken(): string | null {
  if (typeof window === "undefined") return null
  return localStorage.getItem("jwt") || localStorage.getItem("token")
}

export function setToken(token: string) {
  if (typeof window === "undefined") return
  localStorage.setItem("jwt", token)
}

export function logout() {
  if (typeof window === "undefined") return
  localStorage.removeItem("jwt")
  localStorage.removeItem("token")
  localStorage.removeItem("usuario_nome")
}

export function isAuthenticated(): boolean {
  return Boolean(getToken())
}
export function isAdmin(): boolean {
  if (typeof window === "undefined") return false
  const role = localStorage.getItem("usuario_role")
  return role === "admin"
}

