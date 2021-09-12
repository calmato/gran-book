import React, { ReactElement } from 'react';
import { ThemeProvider } from 'react-native-elements';
import { Provider } from 'react-redux';
import { AuthProvider } from '~/context/auth';
import { UiContext } from '~/lib/context';
import { createApplicationInitialState } from '~/lib/context/ui';
import MainRoute from '~/routes';
import store from '~/store';
import { THEME } from '~~/constants/theme';

const App = function App(): ReactElement {
  const [applicationState, setApplicationState] = React.useState(createApplicationInitialState());

  return (
    <Provider store={store}>
      <ThemeProvider theme={THEME}>
        <UiContext.Provider value={{ applicationState, setApplicationState }}>
          <AuthProvider>
            <MainRoute />
          </AuthProvider>
        </UiContext.Provider>
      </ThemeProvider>
    </Provider>
  );
};

export default App;
