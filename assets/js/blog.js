document.addEventListener("DOMContentLoaded", function () {
    const path = window.location.pathname;
    const postId = path.split('/').pop(); // 获取 URL 中的博客 ID

    fetch(`/post/${postId}`)
        .then(response => response.json())
        .then(post => {
            document.getElementById("post-title").innerText = post.title;
            document.getElementById("post-date").innerText = post.date;
            document.getElementById("post-content").innerText = post.content;
        })
        .catch(err => console.log('Error fetching blog post:', err));
});