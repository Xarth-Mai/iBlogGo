document.addEventListener("DOMContentLoaded", function () {
    fetch('/list')
        .then(response => response.json())
        .then(posts => {
            const blogList = document.getElementById("blog-list");
            posts.forEach(post => {
                const li = document.createElement("li");
                li.innerHTML = `<a href="/blog/${post.id}">${post.title}</a> - ${post.date}`;
                blogList.appendChild(li);
            });
        })
        .catch(err => console.log('Error fetching blog list:', err));
});