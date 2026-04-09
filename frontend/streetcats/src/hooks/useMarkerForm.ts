import { useState } from "react"

interface ExtraMarker {
  position: [number, number]
  markerText: string
  id: string
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
  const [notes, setNotes] = useState("")

  const handleMarkerPlaced = (position: [number, number]) => {
    const id = Date.now().toString()
    setExtraMarkers(prev => [...prev, { position, markerText: "Modifica...", id }])
    setPendingMarker({ position, id })
    setSidebarOpen(true)
  }

  const handleSave = () => {
    if (!pendingMarker) return

    const breedText = selectedBreed
      ? breedOptions.find(b => b.value === selectedBreed)?.label ?? selectedBreed
      : "Razza non selezionata"
    const markerText = `${breedText}${notes ? ` - ${notes}` : ""}`

    setExtraMarkers(prev =>
      prev.map(marker =>
        marker.id === pendingMarker.id
          ? { ...marker, markerText }
          : marker
      )
    )
    resetForm()
  }

  const handleCancel = () => {
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
    setNotes("")
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
    notes,
    // Setters
    setMarkerMode,
    setSidebarOpen,
    setSelectedBreed,
    setNotes,
    // Handlers
    handleMarkerPlaced,
    handleSave,
    handleCancel,
    closeSidebar,
  }
}
