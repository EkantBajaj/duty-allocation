import React from 'react';
import Index from './pages/Index';
import {Route,BrowserRouter,Switch,Routes} from "react-router-dom";
import SpotUsers from "./pages/SpotUsers";

const App = () => (
    <div className="App">
        <BrowserRouter>
            <Routes>
            <Route path="/" element={<Index />}/>
                <Route path="/user/:id/:name" element={<SpotUsers />}/>
            </Routes>
        </BrowserRouter>
    </div>
);

export default App;
