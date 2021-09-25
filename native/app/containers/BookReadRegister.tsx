import { RouteProp } from "@react-navigation/native";
import { StackNavigationProp } from "@react-navigation/stack";
import React, { ReactElement, useContext } from "react";
import { AuthContext } from "~/context/auth";
import { BookContext } from "~/context/book";
import BookReadRegister from "~/screens/BookReadRegister";
import { registerOwnBook } from "~/store/usecases/v2/book";
import { ImpressionForm } from "~/types/forms";
import { BookshelfTabStackParamList } from "~/types/navigation";

interface Props {
  route: RouteProp<BookshelfTabStackParamList, "BookReadRegister">;
  navigation: StackNavigationProp<
    BookshelfTabStackParamList,
    "BookReadRegister"
  >;
}

const ConnectedBookReadRegister = function ConnectedBookReadRegister(
  props: Props
): ReactElement {
  const { authState } = useContext(AuthContext);
  const { fetchBooks } = useContext(BookContext);

  const actions = React.useMemo(
    () => ({
      registerReadBookImpression(bookId: number, impression: ImpressionForm) {
        return registerOwnBook(
          {
            userId: authState.id,
            bookId,
            status: "read",
            impressionForm: impression,
          },
          authState.token
        );
      },
      fetchBooks,
    }),
    []
  );

  return <BookReadRegister {...props} actions={actions} />;
};

export default ConnectedBookReadRegister;
