
import { BrowserRouter, Route, Routes } from 'react-router-dom'

export default function Router() {
        return (
                <BrowserRouter>
                        <Nav>
                                <Routes>
                                        <Route path='/' element={<Home />} />
                                        <Route path='*' element={<NotFound />} />
                                </Routes>

                        </Nav>

                </BrowserRouter>

        )
}