document.addEventListener("DOMContentLoaded", function () {
    // post 模板
    const postTemplate = {
        id: 1,
        title: "Default Title",
        content: "Default Content",
        date: "2024-01-01"
    };

    const path = window.location.pathname;
    const postId = path.split('/').pop(); // 获取 URL 中的博客 ID

    fetch(`/post/${postId}`)
        .then(response => response.json())
        .then(post => {
            // 合并模板和获取的数据
            const updatedPost = { ...postTemplate, ...post };

            // 填充页面内容
            document.getElementById("post-title").innerText = updatedPost.title;
            document.getElementById("post-date").innerText = updatedPost.date;
            document.getElementById("post-content").innerText = updatedPost.content;
        })
        .catch(err => console.log('Error fetching blog post:', err));
});
