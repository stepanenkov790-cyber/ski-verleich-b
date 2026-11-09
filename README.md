# Ski-Verleih – Prototyp

Ein einfacher Prototyp für eine Ski-Verleihverwaltung mit Go-Backend und HTML/CSS/JavaScript-Frontend.

## Voraussetzungen
- Go 1.22 oder neuer

## Starten
Im Projektordner ausführen:

```bash
go run .
```

Danach im Browser öffnen: http://localhost:8080

## Tests und Formatierung

```bash
go test ./...
gofmt -w .
```

## Funktionen
- Dashboard mit verfügbaren und ausgeliehenen Ski sowie Personenanzahl
- Ski hinzufügen und löschen
- Personen hinzufügen und löschen
- Ausleihe für einen registrierten Kunden erstellen\n- Flexible Mietpreise pro Ski und Tag\n- Gesamtpreis wird anhand der Mietdauer automatisch berechnet
- Reservierungen erstellen und anzeigen

## API
- `GET /api/dashboard`
- `GET, POST /api/skis`
- `PUT, DELETE /api/skis/{id}`
- `GET, POST /api/persons`
- `PUT, DELETE /api/persons/{id}`
- `GET, POST /api/loans`
- `GET, POST /api/reservations`

## Hinweis
Die Daten werden nur im Arbeitsspeicher gehalten. Nach einem Neustart des Servers werden Änderungen zurückgesetzt. Das ist bewusst so umgesetzt, weil die Aufgabenbeschreibung einen In-Memory-Store für den Prototyp vorsieht.

## Flexible Preisverwaltung (Aufgabenteil b)
- Für jeden Ski wird ein individueller Tagesmietpreis in Euro gespeichert.
- Bei einer Ausleihe wird die Mietdauer in Tagen angegeben.
- Der Gesamtpreis wird automatisch berechnet: Tagespreis × Mietdauer.
- Der Tagespreis wird bei der Ausleihe übernommen, damit spätere Preisänderungen alte Ausleihen nicht verändern.
