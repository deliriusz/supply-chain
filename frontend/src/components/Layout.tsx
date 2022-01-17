import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Container } from "semantic-ui-react";
import Dashboard from "./admin/Dashboard";
import Footer from "./Footer";
import Header from "./Header";
import Main from "./Main";
import ProductDetails from "./ProductDetails";
import ProductsTable from "./ProductsTable";
import ProductStatus from "./ProductStatus";
import Trace from "./Trace";
import 'semantic-ui-css/semantic.min.css'
import './Layout.css'

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