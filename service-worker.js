var cachedFilePaths = [
  '/static/js/helpers.js',
  '/static/js/elasticlunr.min.js',
  '/static/js/search.js',
  '/static/js/main.js',
  '/static/css/style.css',
  '/resources.js'
];

var cacheVersion = 'v1';

self.addEventListener('install', function(installEvt) {
  installEvt.waitUntil(
    caches.open(cacheVersion).then(function(cache) {
      var allCachedFiles = cachedFilePaths.concat('/');
      return cache.addAll(allCachedFiles);
    })
  );
});

function hasFileBeenCached(request, filePaths) {
  var requestUrl = request.url

  if (request.headers.get('accept').indexOf('text/html') > -1) {
    return true;
  }

  for (var i = 0; i < filePaths.length; i++) {
    if (requestUrl.indexOf(cachedFilePaths[i]) > -1) return true;
  }

  return false;
}

self.addEventListener('fetch', function(fetchEvt) {
  var request = fetchEvt.request;
  if (hasFileBeenCached(request, cachedFilePaths)) {
    fetchEvt.respondWith(
      fetch(request).then(function(response) {
        return response.ok ? response : caches.match(request);
      }).catch(function() {
        return caches.match(request);
      })
    );
  }
});
