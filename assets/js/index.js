document.addEventListener("DOMContentLoaded", function () {
    // PostMetadata 模板
    const PostMetadataTemplate = {
        id: 1,
        title: "New Title",
        date: "2023-09-15T00:00:00Z"
    };

    fetch('/list')
        .then(response => response.json())
        .then(posts => {
            const blogList = document.querySelector(".blog-list");

            posts.forEach(item => {
                // 使用修改后的PostMetadata对象
                const updatedPost = { ...PostMetadataTemplate, ...item };

                // 格式化日期
                const postDate = new Date(updatedPost.date); // 将日期字符串转换为 Date 对象
                const formattedDate = postDate.toLocaleDateString('en-US', {
                    weekday: 'long', // 显示星期几
                    year: 'numeric', // 显示年份
                    month: 'long',   // 显示月份（全名）
                    day: 'numeric',  // 显示日期
                });

                // 创建博客卡片元素
                const card = document.createElement("div");
                card.classList.add("blog-card");

                // 填充卡片内容
                card.innerHTML = `
                    <h2>${updatedPost.title}</h2>
                    <p>Date: ${formattedDate}</p>
                    <p><a href="/blog/${updatedPost.id}">Read more...</a></p>
                `;

                blogList.appendChild(card);
            });
        })
        .catch(err => console.log('Error fetching blog list:', err));
});
