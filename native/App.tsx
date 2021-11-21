import React, { ReactElement } from 'react';
import { ThemeProvider } from 'react-native-elements';
import { Provider } from 'react-redux';
import { AuthProvider } from '~/context/auth';
import { UiProvider } from '~/context/ui';
import MainRoute from '~/routes';
import store from '~/store';
import { THEME } from '~~/constants/theme';

const App = function App(): ReactElement {
  return (
    <Provider store={store}>
      <ThemeProvider theme={THEME}>
        <UiProvider>
          <AuthProvider>
            <MainRoute />
          </AuthProvider>
        </UiProvider>
      </ThemeProvider>
    </Provider>
  );
};

export default App;
