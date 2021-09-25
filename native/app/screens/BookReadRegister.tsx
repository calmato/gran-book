import { RouteProp } from "@react-navigation/native";
import { StackNavigationProp } from "@react-navigation/stack";
import dayjs from "dayjs";
import React, { ReactElement, useCallback, useState } from "react";
import { StyleSheet, View, ScrollView, Text, TextInput } from "react-native";
import { Button } from "react-native-elements";
import BookNameAuthorRegister from "~/components/organisms/BookNameAuthorRegister";
import HeaderWithBackButton from "~/components/organisms/HeaderWithBackButton";
import ReadDate from "~/components/organisms/ReadDate";
import { BookshelfV1Response } from "~/types/api/bookshelf_apiv1_response_pb";
import { ImpressionForm } from "~/types/forms";
import { BookshelfTabStackParamList } from "~/types/navigation";
import { COLOR, FONT_SIZE } from "~~/constants/theme";

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  impressionFormLabel: {
    fontSize: FONT_SIZE.TITLE_SUBHEADER,
    paddingLeft: 16,
    marginTop: 10,
    marginBottom: 8,
    fontWeight: "bold",
    color: COLOR.GREY,
  },
  impressionForm: {
    backgroundColor: COLOR.BACKGROUND_WHITE,
    padding: 0,
    fontSize: FONT_SIZE.TEXT_INPUT,
    paddingHorizontal: 16,
    height: 160,
    marginBottom: 16,
  },
});

interface Props {
  route: RouteProp<BookshelfTabStackParamList, "BookReadRegister">;
  navigation: StackNavigationProp<
    BookshelfTabStackParamList,
    "BookReadRegister"
  >;
  actions: {
    registerReadBookImpression: (
      bookId: number,
      impression: ImpressionForm
    ) => Promise<BookshelfV1Response.AsObject | undefined>;
    fetchBooks: () => Promise<void>;
  };
}

const BookReadRegister = function BookReadRegister(props: Props): ReactElement {
  const book = props.route.params.book;
  const { registerReadBookImpression, fetchBooks } = props.actions;

  const [impressionData, setState] = useState({
    date: new Date(),
    impression: "",
    isDateUnknown: false,
  });

  const handleRegisterButtonClick = useCallback(async () => {
    await registerReadBookImpression(book.id, {
      impression: impressionData.impression,
      readOn: dayjs(impressionData.date).format("YYYY-MM-DD"),
    });
    await fetchBooks();

    props.navigation.navigate("Bookshelf");
  }, []);

  return (
    <View style={styles.container}>
      <HeaderWithBackButton
        title="読んだ本登録"
        onPress={() => {
          props.navigation.goBack();
        }}
      />
      <ScrollView>
        <BookNameAuthorRegister
          title={book.title}
          imageUrl={book.thumbnailUrl}
          author={book.author}
        />
        <ReadDate
          date={impressionData.date}
          handleSetDate={(date) => setState({ ...impressionData, date: date })}
          isDateUnknown={impressionData.isDateUnknown}
          handleIsDateUnknown={(isDateUnknown) =>
            setState({ ...impressionData, isDateUnknown: isDateUnknown })
          }
        />
        <Text style={styles.impressionFormLabel}>感想</Text>
        <TextInput
          style={styles.impressionForm}
          placeholder="ここに感想を入力"
          onChangeText={(text) =>
            setState({ ...impressionData, impression: text })
          }
          value={impressionData.impression}
          maxLength={1000}
          multiline={true}
        />
        <View style={{ alignItems: "center", marginBottom: 20 }}>
          <Button onPress={handleRegisterButtonClick} title="本を登録する" />
        </View>
      </ScrollView>
    </View>
  );
};

export default BookReadRegister;
