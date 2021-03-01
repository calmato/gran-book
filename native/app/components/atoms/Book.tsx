import React, { ReactElement } from 'react';
import { StyleSheet, View, Text, Image, ImageURISource, ViewStyle, ImageProps } from 'react-native';
import { Tile } from 'react-native-elements';
import { COLOR } from '~~/constants/theme';

const styles = StyleSheet.create({
  containerStyle: {
    backgroundColor: COLOR.BACKGROUND_WHITE,
    borderRadius: 5,
  },
  contentContainerStyle: {
    height: 60,
    paddingLeft: 4,
    paddingRight: 4,
    paddingTop: 8,
  },
  titleStyle: {
    color: COLOR.TEXT_DEFAULT,
    fontWeight: '900',
    fontSize: 12,
  },
  authorStyle: {
    color: COLOR.TEXT_GRAY,
    fontSize: 8,
  }
});

const imageProps: Partial<ImageProps> = {
  resizeMode: 'contain',
};

interface Props {
  title: string,
  author: string,
  image: ImageURISource,
  width: number,
  height: number,
  onPress?: () => undefined | void,
  style?: ViewStyle,
}

const Book = function Book(props: Props): ReactElement {
  const {title, author, image, width, height, style} = props;

  return (
    <View style={style}>
      <Tile
        containerStyle={styles.containerStyle}
        contentContainerStyle={styles.contentContainerStyle}
        titleStyle={styles.titleStyle}
        imageSrc={image}
        imageProps={imageProps}
        title={title}
        width={width}
        height={height}
        titleNumberOfLines={2}
        onPress={props.onPress}
      >
        <Text
          style={styles.authorStyle}
          numberOfLines={1}
        >
          {author}
        </Text>
      </Tile>
    </View>
  );
};

export default Book;
