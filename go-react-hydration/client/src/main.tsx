import { hydrateRoot } from 'react-dom/client'
import App from './App'
import './index.css'

// Get the product data that was embedded in the HTML
const productDataElement = document.getElementById('product-data')
const productData = productDataElement ? JSON.parse(productDataElement.textContent || '{}') : {}

// Find the root element that already contains the server-rendered HTML
const rootElement = document.getElementById('root')

if (rootElement) {
  // Hydrate the existing HTML with React
  hydrateRoot(
      rootElement, 
      <App initialProduct={productData} />
  )
}