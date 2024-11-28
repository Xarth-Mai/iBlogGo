document.addEventListener("DOMContentLoaded", function () {
    // post 模板
    const postTemplate = {
        id: 0,
        title: "Default Title",
        content: "Default Content",
        date: "2050-10-10T00:00:00+00:00" // 默认时间
    };

    const path = window.location.pathname;
    const postId = path.split('/').pop() || 1; // 获取 URL 中的博客 ID，如果没有则默认为 1

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
            document.getElementById("post-date").innerText = formattedDate;
            document.getElementById("post-content").innerHTML = updatedPost.content;

            // 高亮代码
            hljs.highlightAll();
        })
        .catch(err => console.log('Error fetching blog post:', err));
});
