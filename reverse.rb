# Без использования стандартных функций, написать функцию reverse, которая принимает массив и возвращает массив в обратном порядке

class Array
  def reverse
    result, index, arr_length = [],0,length
    while index < arr_length do
      result << pop
      index += 1
    end
    result
  end
end
# test
p ["a","b","c"].reverse
