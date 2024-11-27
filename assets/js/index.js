document.addEventListener("DOMContentLoaded", function () {
    // post 模板
    const postsTemplate = {
        id: 1,
        title: "New Title",
        date: "2024-01-01"
    };

    fetch('/list')
        .then(response => response.json())
        .then(posts => {
            const blogList = document.querySelector(".blog-list");

            posts.forEach(item => {
                // 使用修改后的post对象
                const updatedPost = {...postsTemplate, ...item};

                // 创建博客卡片元素
                const card = document.createElement("div");
                card.classList.add("blog-card");

                // 填充卡片内容
                card.innerHTML = `
                <h2>${updatedPost.title}</h2>
                <p>Date: ${updatedPost.date}</p>
                <p><a href="/blog/${updatedPost.id}">Read more...</a></p>
                `;

                blogList.appendChild(card);
            });
        })
        .catch(err => console.log('Error fetching blog list:', err));
});
