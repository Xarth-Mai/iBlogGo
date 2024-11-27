document.addEventListener("DOMContentLoaded", function () {
    // post 模板
    const postTemplate = {
        id: 1,
        title: "Default Title",
        content: "Default Content",
        date: "2023-09-15T00:00:00Z" // 默认时间
    };

    const path = window.location.pathname;
    const postId = path.split('/').pop(); // 获取 URL 中的博客 ID

    fetch(`/post/${postId}`)
        .then(response => response.json())
        .then(post => {
            // 合并模板和获取的数据
            const updatedPost = {...postTemplate, ...post};

            // 修改 HTML 标题
            document.title = updatedPost.title;

            // 格式化日期
            const postDate = new Date(updatedPost.date);  // 将日期字符串解析为 Date 对象

            // 使用 toLocaleDateString 格式化日期为易读格式
            const formattedDate = postDate.toLocaleDateString('en-US', {
                weekday: 'long', // 星期几（如 Monday）
                year: 'numeric', // 年份
                month: 'long', // 月份（如 November）
                day: 'numeric', // 日期
            });

            // 填充页面内容
            document.getElementById("post-title").innerText = updatedPost.title;
            document.getElementById("post-date").innerText = formattedDate; // 显示格式化后的日期
            document.getElementById("post-content").innerText = updatedPost.content;
        })
        .catch(err => console.log('Error fetching blog post:', err));
});
