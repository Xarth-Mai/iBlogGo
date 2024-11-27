document.addEventListener("DOMContentLoaded", function () {
    const post = {
        id: 1,
        title: "Title",
        date: "0000-00-00"
    };
    fetch('/list')
        .then(response => response.json())
        .then(posts => {
            const blogList = document.querySelector(".blog-list"); // 选择 .blog-list 容器
            posts.forEach(post => {
                // 创建博客卡片元素
                const card = document.createElement("div");
                card.classList.add("blog-card");

                // 填充卡片内容
                card.innerHTML = `
                <h2>${post.title}</h2>
                <p>Date: ${post.date}</p>
                <p><a href="/blog/${post.id}">Read more...</a></p>
                `;
                
                // 将卡片添加到博客列表中
                blogList.appendChild(card);
            });
        })
        .catch(err => console.log('Error fetching blog list:', err));
});
