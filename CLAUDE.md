# Projektregeln

## Ziel
Nur ein Prototyp. Priorität: schnell verständlich und leicht änderbar.

## Stack
- Backend: Go, Structs und Methoden
- Datenhaltung: In-Memory, keine Datenbank und kein ORM
- Frontend: HTML/CSS/JavaScript
- Versionierung: lokales Git, GitHub bei Bedarf

## Architektur
- `main.go`: Einstiegspunkt, HTTP-Server und Routing
- `internal/domain/`: zentrale Structs
- `internal/store/`: thread-sicherer In-Memory-Store
- `internal/handler/`: HTTP-Handler
- `web/`: Frontend

## Konventionen
- Methoden an Structs, wo sinnvoll
- Interfaces nur bei echtem Bedarf
- Kein Dependency-Injection-Framework
- Keine neuen Dependencies ohne Nachfrage
- Store muss thread-safe sein
- Keine Secrets im Code
- Fehler nicht verschlucken

## Frontend
- Möglichst minimales JavaScript
- Backend-Aufrufe ausschließlich über HTTP-API
- Keine direkten Zugriffe auf In-Memory-Strukturen

## Regeln
- Keine generierten Dateien von Hand editieren
- Bei Unklarheit nachfragen statt raten
- Einfach und klar statt unnötig komplex
