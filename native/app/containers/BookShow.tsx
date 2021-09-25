import { RouteProp } from "@react-navigation/native";
import { StackNavigationProp } from "@react-navigation/stack";
import React, { ReactElement, useContext } from "react";
import { BookContext } from "~/context/book";
import BookShow from "~/screens/BookShow";
import { BookshelfTabStackParamList } from "~/types/navigation";

interface Props {
  route:
    | RouteProp<BookshelfTabStackParamList, "SearchResultBookShow">
    | RouteProp<BookshelfTabStackParamList, "BookShow">;
  navigation:
    | StackNavigationProp<BookshelfTabStackParamList, "SearchResultBookShow">
    | StackNavigationProp<BookshelfTabStackParamList, "BookShow">;
}

export default function ConnectedBookShow(props: Props): ReactElement {
  const { registerBook, fetchBooks } = useContext(BookContext);

  return (
    <BookShow
      route={props.route}
      navigation={props.navigation}
      actions={{ registerBook, fetchBooks }}
    />
  );
}
