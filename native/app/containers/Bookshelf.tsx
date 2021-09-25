import { StackNavigationProp } from "@react-navigation/stack";
import React, { useContext, useEffect } from "react";
import { BookContext } from "~/context/book";
import Bookshelf from "~/screens/Bookshelf";
import { BookshelfTabStackParamList } from "~/types/navigation";

interface Props {
  navigation: StackNavigationProp<BookshelfTabStackParamList, "Bookshelf">;
}

export default function ConnectedBookshelf(props: Props) {
  const { viewBooks, fetchBooks } = useContext(BookContext);

  useEffect(() => {
    const f = async () => {
      await fetchBooks();
    };
    f();
  }, [fetchBooks]);

  return (
    <Bookshelf
      actions={{ fetchBooks }}
      books={viewBooks}
      navigation={props.navigation}
    />
  );
}
