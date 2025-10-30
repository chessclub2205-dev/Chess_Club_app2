self.addEventListener('install', (event) => {
  event.waitUntil(
    caches.open('app-static-v1').then((cache) =>
      cache.addAll(['/', '/index.html'])
    )
  )
})

self.addEventListener('activate', () => self.clients.claim())

self.addEventListener('fetch', (event) => {
  event.respondWith(
    caches.match(event.request).then((cached) => cached || fetch(event.request))
  )
})

