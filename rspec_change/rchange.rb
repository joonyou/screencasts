class Foo
  def repeat
    2.times { repeated }
  end

  def repeated; end
end

describe Foo do
  let(:foo) { stub method1:"hello" }

  context "stub" do
    it "is deprecated to use double" do
      expect(foo.method1).to eq "hello"
    end
  end

  context "stub!" do
    let(:foo) { Foo.new }

    it "is deprecated" do
      foo.stub! bar:"bar"
      expect(foo.bar).to eq "bar"
    end
  end

  context "any_number_of_times" do
    it "is deprecated" do
      Foo.any_instance.should_receive(:repeated).any_number_of_times
      Foo.new.repeat
    end
  end
end

