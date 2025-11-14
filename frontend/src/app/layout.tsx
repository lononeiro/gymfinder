import "./globals.css";
import Header from "@/components/Header";

export const metadata = {
  title: "GymFinder",
  description: "Encontre e avalie academias na sua região.",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="pt-BR">
      <body className="bg-yellow-50 min-h-screen">
        <Header />
        <main className="pt-6">{children}</main>
      </body>
    </html>
  );
}
