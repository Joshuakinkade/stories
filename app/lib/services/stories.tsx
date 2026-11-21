export interface Story {
  id: string
  title: string
  slug: string
  elements: Element[]
  createdAt: Date
  updatedAt: Date
}

export interface Element {
  id: string
  type: string
  contents: object
} 
