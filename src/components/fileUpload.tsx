import React, { useState } from "react";

function FileUploader() {
  const [selectedFile, setSelectedFile] = useState(null);
  const [uploading, setUploading] = useState(false);

  const handleFileChange = (event) => {
    if (event.target.files && event.target.files[0]) {
      setSelectedFile(event.target.files[0]);
    }
  };

  const handleUpload = async () => {
    if (!selectedFile) return;

    const formData = new FormData();
    formData.append("file", selectedFile);

    setUploading(true);
    try {
      const response = await fetch("http://localhost:3000/upload", {
        method: "POST",
        body: formData,
      });

      if (response.ok) {
        alert("Arquivo enviado com sucesso!");
      } else {
        alert("Erro ao enviar o arquivo.");
      }
    } catch (err) {
      alert("Erro na requisição.");
    } finally {
      setUploading(false);
    }
  };

  return (
    <div className="w-full flex items-center justify-between gap-4">
      <input
        type="file"
        id="fileInput"
        onChange={handleFileChange}
        className="hidden"
      />
      <label
        htmlFor="fileInput"
        className="px-4 py-2 bg-red-900 text-white rounded hover:bg-red-950 cursor-pointer"
      >
        Escolher arquivo
      </label>
      {selectedFile && (
        <p className="text-sm text-gray-700">
          Selecionado: {selectedFile.name}
        </p>
      )}
      <button
        onClick={handleUpload}
        disabled={!selectedFile || uploading}
        className={`ml-auto px-4 py-2 rounded text-white ${
          uploading
            ? "bg-gray-400 cursor-not-allowed"
            : "bg-black hover:bg-gray-900"
        }`}
      >
        {uploading ? "Enviando..." : "Enviar"}
      </button>
    </div>
  );
}

export default FileUploader;
