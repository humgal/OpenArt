import React from 'react';
import { View, Text } from "react-native";
import ArtCard from '../common/Card/ArtCard'

const Home = () => {
  return (
    <View >
      <ArtCard creator={'Pawel Czerwinski'} id={'11'} name={'Silent Wave'} saleStatus={0} uploadUrl={''} ></ArtCard>
    </View>
  );
};

export default Home;