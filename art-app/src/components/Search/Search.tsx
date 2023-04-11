import { View, Text } from "react-native";
import SearchInput from "./SearchInput/SearchInput";
import style from "./Search.css"
import SearchOption from "./SearchOption/SearchOption"

const Search = () => {
  return (
    <View style={style.Search}>
      <SearchInput/>
      <SearchOption></SearchOption>
    </View>
  );
};
export default Search;