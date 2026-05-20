# HTTP Client Configuration dla IDEA (IntelliJ)

Dokumentacja do testów manualnych E2E używając wbudowanego HTTP Client w IntelliJ IDEA.

## Setup

1. Otwórz plik `jobs.http` w IntelliJ IDEA
2. IDEA automatycznie rozpozna format HTTP Client
3. Możesz teraz wysyłać żądania bezpośrednio z edytora

## Używanie

### Wysłanie żądania

- Kliknij na zieloną ikonę ▶️ (Run) obok żądania
- Lub użyj skrótu: `Ctrl+Alt+Enter` (Windows/Linux) lub `Cmd+Alt+Enter` (Mac)

### Zmienne środowiskowe

Możesz ustawić zmienne w pliku `http-client.env.json`:

```json
{
  "dev": {
    "baseUrl": "http://localhost:8080",
    "jobId": "1"
  },
  "prod": {
    "baseUrl": "https://api.example.com",
    "jobId": "123"
  }
}
```

W pliku `jobs.http` użyj zmiennych:

```http
GET {{baseUrl}}/idp/jobs/{{jobId}}
```

## Dokumentacja

- [IntelliJ HTTP Client Documentation](https://www.jetbrains.com/help/idea/http-client-in-product-code-editor.html)
- [HTTP Client Reference](https://www.jetbrains.com/help/idea/exploring-http-syntax.html)
