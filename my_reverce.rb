class Array
  def my_reverse
    res,i,l = [],-1,size
    while l >= 1
      res << [i]
      l -= 1
      i -= 1
    end
    res
  end
end

p ["a","b","c"].my_reverse
