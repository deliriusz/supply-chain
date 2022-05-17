import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Container } from "semantic-ui-react";
import 'semantic-ui-css/semantic.min.css'
import './style.css'
import Header from "../Header";
import Dashboard from "../../pages/AdminDashboard";
import Main from "../../pages/Main";
import ProductModelDetails from "../../pages/ProductModelDetails";
import ProductModelsTable from "../../pages/ProductModelsTable";
import ProductStatus from "../../pages/ProductStatus";
import MyPurchases from "../../pages/MyPurchases";
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
                     <Route path="/purchase" element={<MyPurchases />} />
                     <Route path="/purchase/:id" element={<ProductStatus />} />
                     <Route path="/product" element={<ProductModelsTable />} />
                     <Route path="/product/:productId" element={<ProductModelDetails />} />
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