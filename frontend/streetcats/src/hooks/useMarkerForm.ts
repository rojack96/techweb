import { useState } from "react"

interface ExtraMarker {
  position: [number, number]
  id: string
  title: string
  description: string
  breed: string
}

interface PendingMarker {
  position: [number, number]
  id: string
}

export const breedOptions = [
  { value: "siamese", label: "Siamese" },
  { value: "maine_coon", label: "Maine Coon" },
  { value: "persian", label: "Persiano" },
  { value: "ragdoll", label: "Ragdoll" },
  { value: "sphynx", label: "Sphynx" },
]

export function useMarkerForm() {
  const [markerMode, setMarkerMode] = useState(false)
  const [sidebarOpen, setSidebarOpen] = useState(false)
  const [extraMarkers, setExtraMarkers] = useState<ExtraMarker[]>([])
  const [pendingMarker, setPendingMarker] = useState<PendingMarker | null>(null)
  const [selectedBreed, setSelectedBreed] = useState<string | undefined>(undefined)
  const [markerTitle, setMarkerTitle] = useState("")
  const [markerDescription, setMarkerDescription] = useState("")

  const handleMarkerPlaced = (position: [number, number]) => {
    const id = Date.now().toString()
    // Aggiungi subito il marker placeholder sulla mappa
    setExtraMarkers(prev => [...prev, {
      position,
      id,
      title: "Modifica...",
      description: "",
      breed: ""
    }])
    setPendingMarker({ position, id })
    setSidebarOpen(true)
  }

  const handleSave = () => {
    if (!pendingMarker) return

    const breedText = selectedBreed
      ? breedOptions.find(b => b.value === selectedBreed)?.label ?? selectedBreed
      : "Razza non selezionata"

    // Aggiorna il marker esistente con i dati completi
    setExtraMarkers(prev =>
      prev.map(marker =>
        marker.id === pendingMarker.id
          ? {
            ...marker,
            title: markerTitle || "Senza titolo",
            description: markerDescription,
            breed: breedText
          }
          : marker
      )
    )
    resetForm()
  }

  const handleCancel = () => {
    // Rimuovi il marker se l'utente cancella
    if (pendingMarker !== null) {
      setExtraMarkers(prev => prev.filter(marker => marker.id !== pendingMarker.id))
    }
    resetForm()
  }

  const resetForm = () => {
    setPendingMarker(null)
    setSidebarOpen(false)
    setMarkerMode(false)
    setSelectedBreed(undefined)
    setMarkerTitle("")
    setMarkerDescription("")
  }

  const closeSidebar = () => {
    setSidebarOpen(false)
    setMarkerMode(false)
  }

  return {
    // State
    markerMode,
    sidebarOpen,
    extraMarkers,
    pendingMarker,
    selectedBreed,
    markerTitle,
    markerDescription,
    // Setters
    setMarkerMode,
    setSidebarOpen,
    setSelectedBreed,
    setMarkerTitle,
    setMarkerDescription,
    // Handlers
    handleMarkerPlaced,
    handleSave,
    handleCancel,
    closeSidebar,
  }
}
