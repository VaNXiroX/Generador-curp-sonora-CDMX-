package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

// Puerto de operación
const PORT = ":8080"

func main() {
	// Servir archivos estáticos desde el directorio actual
	// Esto buscará automáticamente 'index.html'
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	fmt.Println("[JARVIS] Sistema inicializado.")
	fmt.Printf("[JARVIS] Servidor escuchando en: http://localhost%s\n", PORT)
	fmt.Println("[JARVIS] Sirviendo archivo index.html local...")

	// Abrir navegador automáticamente
	openBrowser("http://localhost" + PORT)

	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal("[JARVIS] Error en el servidor: ", err)
	}
}

// Función auxiliar para abrir el navegador
func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		// No soportado
	}
	if err != nil {
		fmt.Printf("[JARVIS] No se pudo abrir el navegador automáticamente. Por favor visite %s\n", url)
	}
}
