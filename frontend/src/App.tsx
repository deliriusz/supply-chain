import React from 'react';
import './App.css';
import GenericErrorModal from './components/GenericErrorModal';
import AuthContext, { loadState, reducer } from './hooks/AuthContext';
import Layout from './layouts/Layout';

const App = () => {
  const [state, dispatch] = React.useReducer(reducer, loadState())

  return (
    <div className="App">
      <AuthContext.Provider value={{
        auth: state,
        dispatcher: dispatch
      }}>
        <GenericErrorModal open={state.isError} description={state.message} header="Login error" />
        <Layout />
      </AuthContext.Provider>
    </div>
  );
}

export default App;
