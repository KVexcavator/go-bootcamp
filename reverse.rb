# Без использования стандартных функций, написать функцию reverse, которая принимает массив и возвращает массив в обратном порядке

class Array

  def reverse
    result = []
    index = 0
    arr_lenght = self.length

    while index < arr_lenght do
      result << self.pop
      index += 1
    end
    result
  end
end
# test
p ["a","b","c"].reverse
