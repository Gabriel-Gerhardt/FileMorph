import Navbar from "../components/navbar";
import Footer from "../components/footer";

function Convert() {
  return (
    <div className="min-h-screen flex flex-col bg-red-900">
      <main className="flex-grow bg-white">
        <Navbar />
        <div className="max-w-7xl mx-auto p-8 text-center text-gray-700">
          <h2 className="text-3xl font-semibold mb-4">Página de Conversão</h2>
          <p>Aqui você pode implementar o seu conversor.</p>
        </div>
      </main>
      <Footer />
    </div>
  );
}

export default Convert;
