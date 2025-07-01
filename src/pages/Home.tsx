import { useNavigate } from "react-router-dom";
import Navbar from "../components/navbar";
import Footer from "../components/footer";

function Home() {
  const navigate = useNavigate();

  return (
    <div className="min-h-screen flex flex-col bg-red-900">
      <main className="flex-grow bg-white">
        <Navbar />

        {/* Hero Section */}
        <section className="flex flex-col items-center justify-center text-center py-24 px-4">
          <h2 className="text-5xl font-bold mb-6 text-red-900">
            Converta arquivos em segundos
          </h2>
          <p className="text-lg max-w-2xl mb-8">
            O <span className="font-semibold">FileMorph</span> é um conversor de
            arquivos simples, rápido e seguro. Transforme PDFs, DOCX e outros
            formatos.
          </p>
          <button
            className="bg-red-800 text-white px-6 py-3 rounded-lg text-lg hover:bg-red-700 transition"
            onClick={() => navigate("/convert")}
          >
            Converter
          </button>
        </section>

        {/* Features */}
        <section id="features" className="pt-16 pb-8 px-6 bg-white">
          <div className="max-w-5xl mx-auto grid md:grid-cols-3 gap-8 text-center">
            <div>
              <h3 className="text-xl font-semibold mb-2">Vários formatos</h3>
              <p>Suporte para PDF, DOCX, PAGES, etc.</p>
            </div>
            <div>
              <h3 className="text-xl font-semibold mb-2">
                Salvamento de dados
              </h3>
              <p>
                Seus arquivos anteriores ficaram salvos em até duas semanas.
              </p>
            </div>
            <div>
              <h3 className="text-xl font-semibold mb-2">Rápido e leve</h3>
              <p>Conversões instantâneas direto no navegador.</p>
            </div>
          </div>
        </section>
      </main>

      <Footer />
    </div>
  );
}

export default Home;
