import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Container } from "semantic-ui-react";
import 'semantic-ui-css/semantic.min.css'
import './style.css'
import Header from "../Header";
import Dashboard from "../../pages/AdminDashboard";
import Main from "../../pages/Main";
import ProductDetails from "../../pages/ProductDetails";
import ProductsTable from "../../pages/ProductsTable";
import ProductStatus from "../../pages/ProductStatus";
import Trace from "../../pages/Trace";
import Footer from "../Footer";

const Layout = () => {
   return (
      <BrowserRouter>
         <div className="wrapper">
            <Container className="main">
               <Header />
               <div className="routing-content">
                  <Routes>
                     <Route path="/" element={<Main />} />
                     <Route path="/trace" element={<Trace />} />
                     <Route path="/trace/:id" element={<ProductStatus />} />
                     <Route path="/product" element={<ProductsTable />} />
                     <Route path="/product/:productId" element={<ProductDetails />} />
                     <Route path="/admin" element={<Dashboard />} />
                  </Routes>
               </div>
            </Container>
            <Footer />
         </div>
      </BrowserRouter>
   );
};

export default Layout;