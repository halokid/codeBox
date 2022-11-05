import './App.css';
import {BrowserRouter, Link, Route, Routes, useLocation} from "react-router-dom";

// import React from "react";

function App() {
  return <BrowserRouter>
    <Routes>
      <Route path="/" element={<Home/>}/>
      <Route path="/about" element={<About/>}/>
      <Route path="/dashboard" element={<Dashboard/>}/>
      <Route path="*" element={<NotFound/>}/>
    </Routes>
  </BrowserRouter>
}

const Header = () => {
  return <ul>
    <li>
      <Link to="/">
        首页
      </Link>
    </li>
    <li>
      <Link to="/about">
        关于
      </Link>
    </li>
    <li>
      <Link to="/dashboard">
        仪表盘
      </Link>
    </li>

  </ul>
};

// todo: `Home = ()` 表示这个组件不用传变量进去
const Home = () => {
  return <>
    <Header/>
    <div>hello world</div>
  </>
};

const About = () => {
  // todo: 使用hook， 获取当前页路径
  const location = useLocation();
  const {from, pathname} = location;

  return <>
    <Header/>
    <div>About what? you are in path {pathname}, and from page {from}</div>
  </>
};

const Dashboard = () => {
  return <>
    <Header/>
    <div>View Dashboard</div>
  </>
};

const NotFound = () => {
  return <>
    <Header/>
    <div>404 Not Found!</div>
  </>
};

export default App;




