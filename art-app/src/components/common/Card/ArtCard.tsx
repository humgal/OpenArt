import * as React from "react";
import { Avatar, IconButton, Card, Text } from "react-native-paper";
import { Item } from "../../../gql/types";

const ArtCard = (props: Item) => {
  return (
    <Card style={{ padding: 5,margin:8 }}>
      <Card.Cover
        style={{ width: "100%" }}
        source={{ uri: "https://picsum.photos/700" }}
      />
      <Card.Content>
        <Text variant="titleLarge">{props.name}</Text>
      </Card.Content>
      <Card.Title
        title={props.creator}
        subtitle={props.creator}
        left={(props) =>  <Avatar.Image size={42} source={require('./2.png')} />}
        right={(props) => <IconButton {...props} icon="cards-heart-outline" onPress={() => {}} />}
      />
    </Card>
  );
};

export default ArtCard;
