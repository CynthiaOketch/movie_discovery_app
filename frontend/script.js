const input = document.getElementById('searchInput');
const results = document.getElementById('results');
let timer;

input.addEventListener('input', () => {
  clearTimeout(timer);
  timer = setTimeout(() => {
    fetch(`/api/search?query=${input.value}`)
      .then(res => res.json())
      .then(data => {
        results.innerHTML = '';
        data.results.forEach(item => {
          const div = document.createElement('div');
          div.innerHTML = `
            <img src="https://image.tmdb.org/t/p/w200${item.poster_path}" alt="${item.title || item.name}" />
            <p>${item.title || item.name}</p>
          `;
          div.addEventListener('click', () => {
            localStorage.setItem('selectedMovieId', item.id);
            window.location.href = 'detail.html';
          });
          results.appendChild(div);
        });
      });
  }, 500);
});
