import React from 'react';
import { View, Text } from "react-native";
import ArtCard from '../common/Card/ArtCard'
import CreatorCard from '../common/Card/CreatorCard';

const Home = () => {
  return (
    <View >
      <ArtCard creator={'Pawel Czerwinski'} id={'11'} name={'Silent Wave'} saleStatus={0} uploadUrl={''} createorId={''} ></ArtCard>
      <CreatorCard id={'1'} followerNum={2} realname={'Kennedy Yanko'} username={'Kennedy Yanko'} img={'https://picsum.photos/700'} bio={' Her methods reflect a dual abstract expressionist-surr…'}/>

    </View>
  );
};

export default Home;