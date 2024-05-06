// import module
const express = require("express");

// PORT
const port = 3000;

// 宣告 router engine
const app = express();

// 宣告 API 路由
app.get("/", (req, res) => {
  res.send("Hello World!");
});

// router engine 監聽 PORT
app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
