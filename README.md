# Pokémon CLI Pokedex

Ein interaktives Command-Line Interface (CLI), um Pokémon zu erkunden, zu fangen und deine eigene Pokedex zu verwalten.

---

## Installation

Stelle sicher, dass du Go installiert hast. Dann kannst du das Projekt direkt starten:

```bash
go run . | tee repl.log
````

Der `tee repl.log` Teil speichert die komplette Session zusätzlich in eine Log-Datei.

---

## Verfügbare Commands

| Command   | Beschreibung                                                                  |
| --------- | ----------------------------------------------------------------------------- |
| `exit`    | Beendet die Pokedex                                                           |
| `help`    | Zeigt diese Hilfeseite an                                                     |
| `map`     | Zeigt die Namen von 20 Locations in der Pokémon-Welt an                       |
| `mapb`    | Zeigt die vorherigen 20 Locations an                                          |
| `explore` | Listet alle Pokémon auf, die an der aktuellen Location gefunden werden können |
| `catch`   | Fängt ein Pokémon und fügt es deiner Pokedex hinzu                            |
| `inspect` | Zeigt Details eines Pokémon an (Name, Größe, Gewicht, Stats, Typen)           |
| `pokedex` | Zeigt eine Liste aller Pokémon, die du bisher gefangen hast                   |

---

## Beispiel

```text
> map
- Pallet Town
- Viridian City
...
> explore
- Bulbasaur
- Charmander
...
> catch Bulbasaur
You caught Bulbasaur!
> inspect Bulbasaur
Name: Bulbasaur
Height: 7
Weight: 69
Stats:
  - hp: 45
  - attack: 49
  - defense: 49
  - special-attack: 65
  - special-defense: 65
  - speed: 45
Types:
  - grass
  - poison
> pokedex
- Bulbasaur
```

---

## Hinweise

* Die Pokedex speichert nur Pokémon, die während der aktuellen Session gefangen wurden.
* `map` und `mapb` helfen dir, die Welt zu erkunden und Pokémon gezielt zu finden.
* `inspect` ist nützlich, um die Stärken und Typen eines Pokémon zu überprüfen, bevor du es fängst.

---

Viel Spaß beim Erkunden und Fangen deiner Pokémon!

